pub mod connection;
pub mod error;

/// From [https://coralogix.com/docs/coralogix-domain/]()
pub enum CoralogixRegion {
    US1,
    US2,
    EU1,
    EU2,
    AP1,
    AP2,
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
        }
        .to_string()
    }
}

pub struct Team {
    /// Team name as it appears in the Coralogix UI
    pub name: String,

    /// Coralogix region where the team is located
    pub region: CoralogixRegion,
}

impl Team {
    /// Create a new Coralogix team
    pub fn new(name: String, region: CoralogixRegion) -> Self {
        Self { name, region }
    }
}
