use crate::Team;

pub struct CoralogixConnection {
    /// Coralogix API key
    pub api_key: String,

    /// Coralogix team
    pub team: Team,
}

impl CoralogixConnection {
    /// Create a new Coralogix client
    pub fn new(api_key: String, team: Team) -> Self {
        Self { api_key, team }
    }
}

