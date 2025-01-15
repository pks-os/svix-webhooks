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
                pub struct TemplateOut {
                        #[serde(rename = "createdAt")]
                        pub created_at: String,
                        #[serde(rename = "description")]
                        pub description: String,
                        #[serde(rename = "featureFlag", skip_serializing_if = "Option::is_none")]
                        pub feature_flag: Option<String>,
                        #[serde(rename = "filterTypes", skip_serializing_if = "Option::is_none")]
                        pub filter_types: Option<Vec<String>>,
                        #[serde(rename = "id")]
                        pub id: String,
                        #[serde(rename = "instructions")]
                        pub instructions: String,
                        #[serde(rename = "instructionsLink", skip_serializing_if = "Option::is_none")]
                        pub instructions_link: Option<String>,
                        #[serde(rename = "kind")]
                        pub kind: models::TransformationTemplateKind,
                        #[serde(rename = "logo")]
                        pub logo: String,
                        #[serde(rename = "name")]
                        pub name: String,
                        #[serde(rename = "orgId")]
                        pub org_id: String,
                        #[serde(rename = "transformation")]
                        pub transformation: String,
                        #[serde(rename = "updatedAt")]
                        pub updated_at: String,
                    }

                    impl TemplateOut {
                    pub fn new(created_at: String, description: String, id: String, instructions: String, kind: models::TransformationTemplateKind, logo: String, name: String, org_id: String, transformation: String, updated_at: String) -> TemplateOut {
                TemplateOut {
                    created_at,
                    description,
                    feature_flag: None,
                    filter_types: None,
                    id,
                    instructions,
                    instructions_link: None,
                    kind,
                    logo,
                    name,
                    org_id,
                    transformation,
                    updated_at,
                    }
                    }
                    }
