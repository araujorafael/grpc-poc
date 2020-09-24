use tonic::{transport::Server, Request, Response, Status};

pub mod discount_grpc {
    tonic::include_proto!("discount");
}

use discount_grpc::{
    server::{Discount, DiscountServer},
    SearchRequest, SearchResponse,
};

use discount;

pub struct DiscountGrpc {}

#[tonic::async_trait]
impl Discount for DiscountGrpc {
    async fn calculate(
        &self,
        request: Request<SearchRequest>, // Accept request of type SearcRequest
    ) -> Result<Response<SearchResponse>, Status> {
        // Return an instance of type SearchResponse
        let req_message = request.into_inner();
        println!("Got a request: {:?}", req_message);

        let product = req_message.product.unwrap();
        let user = req_message.user.unwrap();

        let user_discount = discount::UserDiscount {
            user_id: user.user_id,
            date_of_birth: user.date_of_birth,
            price: product.price,
        };

        let product_discount = match discount::calculate_discount(user_discount) {
            Ok(result) => result,
            Err(err) => {
                println!(">>>> ERROR: {}", err);
                return Err(Status::new(err.code, err.message));
            }
        };

        let reply = discount_grpc::SearchResponse {
            discount_rate: product_discount.discount_rate,
            new_price: product_discount.new_price,
        };

        Ok(Response::new(reply)) // Send back our formatted discount info
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let host = "0.0.0.0";
    let port = "3003";

    let addr = format!("{}:{}", host, port).parse()?;
    let discount = DiscountGrpc {};

    println!("Listening requests on: {}", addr);

    Server::builder()
        .add_service(DiscountServer::new(discount))
        .serve(addr)
        .await?;

    Ok(())
}
