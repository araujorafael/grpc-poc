fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_client(false)
        .compile(&["../../protos/discount.proto"], &["../../protos"])?;
    Ok(())
}
