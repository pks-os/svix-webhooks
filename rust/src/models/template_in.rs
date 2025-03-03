// this file is @generated
use serde::{Deserialize, Serialize};

use super::transformation_template_kind::TransformationTemplateKind;

#[derive(Clone, Debug, Default, PartialEq, Deserialize, Serialize)]
pub struct TemplateIn {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,

    #[serde(rename = "featureFlag")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub feature_flag: Option<String>,

    #[serde(rename = "filterTypes")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub filter_types: Option<Vec<String>>,

    #[serde(skip_serializing_if = "Option::is_none")]
    pub instructions: Option<String>,

    #[serde(rename = "instructionsLink")]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub instructions_link: Option<String>,

    #[serde(skip_serializing_if = "Option::is_none")]
    pub kind: Option<TransformationTemplateKind>,

    pub logo: String,

    pub name: String,

    pub transformation: String,
}

impl TemplateIn {
    pub fn new(logo: String, name: String, transformation: String) -> Self {
        Self {
            description: None,
            feature_flag: None,
            filter_types: None,
            instructions: None,
            instructions_link: None,
            kind: None,
            logo,
            name,
            transformation,
        }
    }
}
