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
        
                /// 
                #[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
                pub enum BackgroundTaskStatus {
                        #[serde(rename = "running")]
                        Running,
                                            #[serde(rename = "finished")]
                        Finished,
                                            #[serde(rename = "failed")]
                        Failed,
                    
                }

                impl std::fmt::Display for BackgroundTaskStatus {
                fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
                match self {
                        Self::Running => write!(f, "running"),
                        Self::Finished => write!(f, "finished"),
                        Self::Failed => write!(f, "failed"),
                }
                }
                }

            impl Default for BackgroundTaskStatus {
            fn default() -> BackgroundTaskStatus {
                Self::Running
            }
            }

