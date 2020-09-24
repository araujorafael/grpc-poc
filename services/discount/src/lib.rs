use chrono::prelude::*;
use std::fmt;
use std::time::{Duration, UNIX_EPOCH};
use tonic::Code;

mod discount_types;

pub struct ServerError {
    pub code: Code,
    pub message: String,
}

// Implement std::fmt::Debug for ServerError
impl fmt::Display for ServerError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "An Error Occurred, {}", self.message) // user-facing output
    }
}

// Implement std::fmt::Debug for ServerError
impl fmt::Debug for ServerError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(
            f,
            "{{ file: {}, line: {}, message: {} }}",
            file!(),
            line!(),
            self.message
        )
    }
}

pub struct UserDiscount {
    pub user_id: String,
    pub date_of_birth: u64,
    pub price: i32,
}

pub struct DiscountProduct {
    pub value: i32,
    pub discount_rate: i32,
    pub new_price: i32,
}

pub fn calculate_discount(user_discount: UserDiscount) -> Result<DiscountProduct, ServerError> {
    println!("Calculating discounts for user: {}", user_discount.user_id);

    let price = price_to_float(user_discount.price);
    let rate = get_user_discount(user_discount.date_of_birth)?;

    let discount_product = DiscountProduct {
        value: user_discount.price,
        discount_rate: float_to_price(rate),
        new_price: float_to_price(price * (1.0 - rate)),
    };

    Ok(discount_product)
}

fn price_to_float(value: i32) -> f32 {
    let float_value = value as f32;
    float_value / 100.0
}

fn float_to_price(value: f32) -> i32 {
    (value * 100.0) as i32
}

fn get_user_discount(birthday: u64) -> Result<f32, ServerError> {
    let unix_date = UNIX_EPOCH + Duration::from_secs(birthday);
    let birthday = DateTime::<Utc>::from(unix_date);
    let now = Local::today();

    let rate = if now.day() == 25 && now.month() == 11 {
        discount_types::Types::BlackFriday
    } else if now.day() == birthday.day() && now.month() == birthday.month() {
        discount_types::Types::Birthday
    } else {
        return Err(ServerError {
            message: "Not elegible for discount".to_string(),
            code: Code::NotFound,
        });
    };

    Ok(discount_types::discount_rates(rate))
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_birthday_discount() {
        let today = Local::now();
        let resp = get_user_discount(today.timestamp() as u64).unwrap();
        assert_eq!(resp, 0.05);
    }

    #[test]
    fn test_discount_error() {
        let date = Local.ymd(2014, 7, 8).and_hms(9, 10, 11);
        match get_user_discount(date.timestamp() as u64) {
            Ok(_) => assert!(false, "Should return error"),
            Err(_) => assert!(true),
        }
    }
}
