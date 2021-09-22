use anyhow::Result;
use crate::db::{DB, User, Merchant};
use rustyline::{error::ReadlineError, Editor};
use std::iter::FromIterator;
use structopt::{clap::AppSettings, StructOpt};

mod db;

#[derive(StructOpt, Debug)]
#[structopt(no_version, global_settings = &[AppSettings::DisableVersion])]
enum NewCmd {
    /// Create a User
    #[structopt(name="user")]
    UserCmd {
        name: String,
        email: String,
        credit_limit: f64,
    },
    /// Create a Merchant
    #[structopt(name="merchant")]
    MerchantCmd {
        name: String,
        discount_percentage: f64,
    },
    /// Create a Transaction
    #[structopt(name="txn")]
    TxnCmd {
        user_name: String,
        merchant_name: String,
        amount: f64,
    },
}

#[derive(StructOpt, Debug)]
#[structopt(no_version, global_settings = &[AppSettings::DisableVersion])]
enum UpdateCmd {
    /// Update a User
    #[structopt(name="user")]
    UserCmd {
        name: String,
        email: String,
        credit_limit: f64,
    },
    /// Update a Merchant
    #[structopt(name="merchant")]
    MerchantCmd {
        name: String,
        discount_percentage: f64,
    },
}

#[derive(StructOpt, Debug)]
#[structopt(no_version, global_settings = &[AppSettings::DisableVersion])]
enum ReportCmd {
    /// Report total discount amount of a merchant
    #[structopt(name="discount")]
    DiscountCmd { merchant_name: String },
    /// Report total due amount of a user
    #[structopt(name="dues")]
    DuesCmd { user_name: String },
    /// Report a list of all users who have reached their credit limit
    #[structopt(name="users-at-credit-limit")]
    UsersAtCreditLimitCmd,
    /// Report all users who have dues and total due amount for all
    #[structopt(name="total-dues")]
    TotalDuesCmd,
}

#[derive(StructOpt, Debug)]
#[structopt(no_version, global_settings = &[AppSettings::DisableVersion])]
struct PaybackCmd {
    user_name: String,
    amount: f64
}

#[derive(StructOpt, Debug)]
#[structopt(no_version, global_settings = &[AppSettings::DisableVersion])]
enum Opt {
    /// Create a new user | merchant | transaction
    New(NewCmd),
    /// Update a user | merchant
    Update(UpdateCmd),
    /// Report discount | dues | users-at-credit-limit | total-dues
    Report(ReportCmd),
    /// Pay back amount for a user
    Payback(PaybackCmd),
}

async fn process_cmd(opt: Opt, db: &DB) -> Result<()> {
    match opt {
        Opt::New(newcmd) => {
            match newcmd {
                NewCmd::UserCmd { name, email, credit_limit } => {
                    let user = db.create_user(User::new(name, email, credit_limit)).await?;
                    println!("{}", user);
                }
                NewCmd::MerchantCmd { name, discount_percentage } => {
                    let merchant = db.create_merchant(Merchant::new(name, discount_percentage)).await?;
                    println!("{}", merchant);
                }
                NewCmd::TxnCmd { user_name, merchant_name, amount } => {
                    let user = db.find_user(user_name).await?;
                    let merchant = db.find_merchant(merchant_name).await?;
                    match db.create_transaction(user, merchant, amount).await {
                        Ok(_) => println!("success!"),
                        Err(e) => println!("rejected! (reason: {})", e)
                    };
                }
            }
        }
        Opt::Update(updatecmd) => {
            match updatecmd {
                UpdateCmd::UserCmd { name, email, credit_limit } => {
                    let user = User::new(name, email, credit_limit);
                    db.update_user(user.clone()).await?;
                    println!("{}", user);
                }
                UpdateCmd::MerchantCmd { name, discount_percentage } => {
                    let merchant = Merchant::new(name, discount_percentage);
                    db.update_merchant(merchant.clone()).await?;
                    println!("{}", merchant);
                }
            }
        }
        Opt::Report(reportcmd) => {
            match reportcmd {
                ReportCmd::DiscountCmd { merchant_name } => {
                    let discount = db.report_discount(merchant_name).await?;
                    println!("{}", discount);
                }
                ReportCmd::DuesCmd { user_name } => {
                    let user_dues = db.report_user_dues(user_name).await?;
                    println!("{}", user_dues);
                }
                ReportCmd::UsersAtCreditLimitCmd => {
                    let users = db.report_users_at_credit_limit().await?;
                    if users.len() > 0 {
                        for user_name in users.iter() {
                            println!("{}", user_name);
                        }
                    } else {
                        println!("No users!");
                    }
                }
                ReportCmd::TotalDuesCmd => {
                    let users = db.report_users_total_dues().await?;
                    if users.len() > 0 {
                        for (user_name, total_dues) in users.iter() {
                            println!("{}: {}", user_name, total_dues);
                        }
                    } else {
                        println!("No users!");
                    }
                }
            }
        }
        Opt::Payback(PaybackCmd { user_name, amount }) => {
            let user = db.find_user(user_name.clone()).await?;
            db.create_payback(user, amount).await?;
            let user_dues = db.report_user_dues(user_name.clone()).await?;
            println!("{}(dues: {})", user_name, user_dues);
        }
    };
    Ok(())
}

#[tokio::main]
async fn main() -> Result<()> {
    let mut rl = Editor::<()>::new();

    if rl.load_history("history.txt").is_err() {
        println!("No previous history");
    }

    let db = DB::new().await?;

    loop {
        let readline = rl.readline("> ");
        match readline {
            Ok(line) => {
                let cmd = format!("pay-later {}", line);
                match Opt::from_iter_safe(Vec::from_iter(cmd.split(" ").map(String::from))) {
                    Ok(cmd) => {
                        match process_cmd(cmd, &db).await {
                            Ok(_) => {},
                            Err(e) => eprintln!("{}", e),
                        };
                    },
                    Err(err) => eprintln!("{}", err),
                }
                rl.add_history_entry(line.as_str());
            }
            Err(ReadlineError::Interrupted) => {
                println!("CTRL-C");
                break;
            }
            Err(ReadlineError::Eof) => {
                println!("CTRL-D");
                break;
            }
            Err(err) => {
                println!("{}", err);
                break;
            }
        }
    }
    rl.append_history("history.txt").unwrap();
    Ok(())
}
