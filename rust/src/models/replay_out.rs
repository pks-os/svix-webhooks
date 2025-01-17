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
pub struct ReplayOut {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "status")]
    pub status: models::BackgroundTaskStatus,
    #[serde(rename = "task")]
    pub task: models::BackgroundTaskType,
}

impl ReplayOut {
    pub fn new(
        id: String,
        status: models::BackgroundTaskStatus,
        task: models::BackgroundTaskType,
    ) -> ReplayOut {
        ReplayOut { id, status, task }
    }
}
