
pub mod proto {
    pub mod resource_manager {
        tonic::include_proto!("coralogix_management_api_grpc");
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
