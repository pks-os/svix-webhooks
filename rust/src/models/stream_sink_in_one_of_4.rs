/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 * 
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
        
                #[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
                pub struct StreamSinkInOneOf4 {
                        #[serde(rename = "config")]
                        pub config: Box<models::SnowflakeConfig>,
                        #[serde(rename = "type")]
                        pub r#type: Type,
                    }

                    impl StreamSinkInOneOf4 {
                    pub fn new(config: models::SnowflakeConfig, r#type: Type) -> StreamSinkInOneOf4 {
                StreamSinkInOneOf4 {
                    config: Box::new(config),
                    r#type,
                    }
                    }
                    }
                    /// 
                    #[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
                    pub enum Type {
                            #[serde(rename = "snowflake")]
                        Snowflake,
                    }

                    impl Default for Type {
                    fn default() -> Type {
                        Self::Snowflake
                    }
                    }
