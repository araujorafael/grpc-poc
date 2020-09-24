pub enum Types {
    Birthday,
    BlackFriday,
}

pub fn discount_rates(discount_type: Types) -> f32 {
    match discount_type {
        Types::Birthday => 0.05,
        Types::BlackFriday => 0.1,
    }
}
