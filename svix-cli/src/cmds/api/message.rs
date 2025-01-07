use chrono::{DateTime, Utc};
use clap::{Args, Subcommand};
use svix::api::*;

use crate::{cli_types::PostOptions, json::JsonOf};

#[derive(Args, Clone)]
pub struct MessageListOptions {
    /// Limit the number of returned items
    #[arg(long)]
    pub limit: Option<i32>,
    /// The iterator returned from a prior invocation
    #[arg(long)]
    pub iterator: Option<String>,
    /// Filter response based on the channel.
    #[arg(long)]
    pub channel: Option<String>,
    /// Only include items created before a certain date.
    #[arg(long)]
    pub before: Option<DateTime<Utc>>,
    /// Only include items created after a certain date.
    #[arg(long)]
    pub after: Option<DateTime<Utc>>,
    /// When `true` message payloads are included in the response.
    #[arg(long)]
    pub with_content: Option<bool>,
    /// Filter messages matching the provided tag.
    #[arg(long)]
    pub tag: Option<String>,
    /// Filter response based on the event type
    #[arg(long)]
    pub event_types: Option<Vec<String>>,
}

impl From<MessageListOptions> for svix::api::MessageListOptions {
    fn from(
        MessageListOptions {
            limit,
            iterator,
            channel,
            before,
            after,
            with_content,
            tag,
            event_types,
        }: MessageListOptions,
    ) -> Self {
        Self {
            limit,
            iterator,
            channel,
            before: before.map(|dt| dt.to_rfc3339()),
            after: after.map(|dt| dt.to_rfc3339()),
            with_content,
            tag,
            event_types,
        }
    }
}

#[derive(Args)]
#[command(args_conflicts_with_subcommands = true)]
#[command(flatten_help = true)]
pub struct MessageArgs {
    #[command(subcommand)]
    pub command: MessageCommands,
}

#[derive(Subcommand)]
pub enum MessageCommands {
    /// List all of the application's messages.
    ///
    /// The `before` and `after` parameters let you filter all items created before or after a certain date. These can be used alongside an iterator to paginate over results
    /// within a certain window.
    ///
    /// Note that by default this endpoint is limited to retrieving 90 days' worth of data
    /// relative to now or, if an iterator is provided, 90 days before/after the time indicated
    /// by the iterator ID. If you require data beyond those time ranges, you will need to explicitly
    /// set the `before` or `after` parameter as appropriate.
    List {
        app_id: String,

        #[clap(flatten)]
        options: MessageListOptions,
    },
    /// Creates a new message and dispatches it to all of the application's endpoints.
    ///
    /// The `eventId` is an optional custom unique ID. It's verified to be unique only up to a day, after that no verification will be made.
    /// If a message with the same `eventId` already exists for the application, a 409 conflict error will be returned.
    ///
    /// The `eventType` indicates the type and schema of the event. All messages of a certain `eventType` are expected to have the same schema. Endpoints can choose to only listen to specific event types.
    /// Messages can also have `channels`, which similar to event types let endpoints filter by them. Unlike event types, messages can have multiple channels, and channels don't imply a specific message content or schema.
    ///
    /// The `payload` property is the webhook's body (the actual webhook message). Svix supports payload sizes of up to ~350kb, though it's generally a good idea to keep webhook payloads small, probably no larger than 40kb.
    Create {
        app_id: String,
        message_in: JsonOf<MessageIn>,
        #[clap(flatten)]
        post_options: Option<PostOptions>,
    },
    /// Get a message by its ID or eventID.
    Get { app_id: String, id: String },
    /// Delete the given message's payload.
    ///
    /// Useful in cases when a message was accidentally sent with sensitive content.
    /// The message can't be replayed or resent once its payload has been deleted or expired.
    ExpungeContent { app_id: String, id: String },
}

impl MessageCommands {
    pub async fn exec(
        self,
        client: &Svix,
        color_mode: colored_json::ColorMode,
    ) -> anyhow::Result<()> {
        match self {
            Self::List { app_id, options } => {
                let resp = client.message().list(app_id, Some(options.into())).await?;
                crate::json::print_json_output(&resp, color_mode)?;
            }
            Self::Create {
                app_id,
                message_in,
                post_options,
            } => {
                let resp = client
                    .message()
                    .create(
                        app_id,
                        message_in.into_inner(),
                        post_options.map(Into::into),
                    )
                    .await?;
                crate::json::print_json_output(&resp, color_mode)?;
            }
            Self::Get { app_id, id } => {
                let resp = client.message().get(app_id, id).await?;
                crate::json::print_json_output(&resp, color_mode)?;
            }
            Self::ExpungeContent { app_id, id } => {
                client.message().expunge_content(app_id, id).await?;
            }
        }

        Ok(())
    }
}
