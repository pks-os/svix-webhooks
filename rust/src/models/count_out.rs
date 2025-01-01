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
                pub struct CountOut {
                        /// There's a ceiling to how many attempts we count. When the limit is reached, this will be `true` to indicate the actual count is higher than given.
                        #[serde(rename = "approximated")]
                        pub approximated: bool,
                        /// The count of attempts matching the query.
                        #[serde(rename = "count")]
                        pub count: i32,
                    }

                    impl CountOut {
                    pub fn new(approximated: bool, count: i32) -> CountOut {
                CountOut {
                    approximated,
                    count,
                    }
                    }
                    }

