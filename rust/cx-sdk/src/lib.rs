use auth::{ApiKey, AuthData};
use com::coralogix::enrichment::v1::{
    AddEnrichmentsRequest, CreateCustomEnrichmentRequest, DeleteCustomEnrichmentRequest,
    GetCustomEnrichmentRequest, GetEnrichmentsRequest, RemoveEnrichmentsRequest,
    UpdateCustomEnrichmentRequest,
};
use cx_api::proto::*;
use error::SdkError;
use std::{fmt::Debug, str::FromStr};
use tonic::{
    transport::{Channel, Endpoint},
    Request,
};

pub mod auth;
pub mod client;
pub mod error;
mod util;

/// From [https://coralogix.com/docs/coralogix-domain/]()
#[derive(Debug, Clone)]
pub enum CoralogixRegion {
    US1,
    US2,
    EU1,
    EU2,
    AP1,
    AP2,
    Custom(String),
}

impl CoralogixRegion {
    /// Endpoint for the corresponding Coralogix region
    /// [https://coralogix.com/docs/coralogix-domain/]()
    pub fn endpoint(&self) -> String {
        match self {
            CoralogixRegion::US1 => "coralogix.us",
            CoralogixRegion::US2 => "cx498.coralogix.com",
            CoralogixRegion::EU1 => "coralogix.com",
            CoralogixRegion::EU2 => "eu2.coralogix.com",
            CoralogixRegion::AP1 => "coralogix.in",
            CoralogixRegion::AP2 => "coralogixsg.com",
            CoralogixRegion::Custom(custom) => custom,
        }
        .to_string()
    }

    pub fn from_env() -> Result<Self, SdkError> {
        let region = std::env::var("CORALOGIX_REGION")?;
        Ok(Self::from_str(&region)?)
    }
}

impl FromStr for CoralogixRegion {
    type Err = SdkError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "us1" => Ok(CoralogixRegion::US1),
            "us2" => Ok(CoralogixRegion::US2),
            "eu1" => Ok(CoralogixRegion::EU1),
            "eu2" => Ok(CoralogixRegion::EU2),
            "ap1" => Ok(CoralogixRegion::AP1),
            "ap2" => Ok(CoralogixRegion::AP2),
            custom => Ok(CoralogixRegion::Custom(custom.to_string())),
        }
    }
}

pub struct CoralogixSdk {
    /// Coralogix API key
    auth: AuthData,

    /// Coralogix team
    pub region: CoralogixRegion,
}

impl CoralogixSdk {
    /// Create a new Coralogix client
    pub fn new(auth: ApiKey, region: CoralogixRegion) -> Self {
        Self {
            auth: (&auth).into(),
            region,
        }
    }

    pub fn connect<T: WrappedClient>(&self) -> Result<T, SdkError> {
        let channel: Channel = Endpoint::from_str(&self.region.endpoint())?.connect_lazy();
        Ok(T::connect(channel))
    }

    pub fn auth_data(&self) -> &AuthData {
        &self.auth
    }
}

pub trait WrappedClient {
    /// Create a new client
    fn connect(channel: Channel) -> Self;
}

pub trait AuthenticatedRequest
where
    Self: Sized,
{
    fn authenticate(self, authentication: &AuthData) -> Request<Self> {
        let auth_metadata = authentication.to_metadata_map();
        let mut req = Request::new(self);
        let metadata = req.metadata_mut();
        *metadata = auth_metadata.clone();
        req
    }
}

impl AuthenticatedRequest for CreateCustomEnrichmentRequest {}
impl AuthenticatedRequest for DeleteCustomEnrichmentRequest {}
impl AuthenticatedRequest for UpdateCustomEnrichmentRequest {}
impl AuthenticatedRequest for GetEnrichmentsRequest {}
impl AuthenticatedRequest for AddEnrichmentsRequest {}
impl AuthenticatedRequest for RemoveEnrichmentsRequest {}
impl AuthenticatedRequest for GetCustomEnrichmentRequest {}
