use anyhow::Context;
use askama::Template;
use axum::{
    http::{HeaderMap, StatusCode},
    response::{Html, IntoResponse, Redirect, Response},
    routing::get,
    Router,
};
use tower_http::services::ServeDir;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let assets_path = std::env::current_dir().unwrap();
    let port = 8000_u16;
    let addr = std::net::SocketAddr::from(([0, 0, 0, 0], port));
    let router = Router::new()
        .route("/", get(|| async { Redirect::permanent("/home") }))
        .route("/home", get(home))
        .route("/about", get(about))
        .route("/blog", get(blog))
        .nest_service(
            "/assets",
            ServeDir::new(format!("{}/assets", assets_path.to_str().unwrap())),
        );

    println!("Router initialized. Listening on port: {}", port);

    axum::Server::bind(&addr)
        .serve(router.into_make_service())
        .await
        .context("error while starting server")?;

    Ok(())
}

async fn home(headers: HeaderMap) -> Response {
    if headers.contains_key("Hx-Request") {
        let template = PartialHomeTemplate {};
        HtmlTemplate(template).into_response()
    } else {
        let template = HomeTemplate {};
        HtmlTemplate(template).into_response()
    }

}

async fn about(headers: HeaderMap) -> Response {
    if headers.contains_key("Hx-Request") {
        let template = PartialAboutTemplate {};
        HtmlTemplate(template).into_response()
    } else {
        let template = AboutTemplate {};
        HtmlTemplate(template).into_response()
    }

}

async fn blog(headers: HeaderMap) -> Response {
    if headers.contains_key("Hx-Request") {
        let template = PartialBlogTemplate {};
        HtmlTemplate(template).into_response()
    } else {
        let template = BlogTemplate {};
        HtmlTemplate(template).into_response()
    }

}

#[derive(Template)]
#[template(path = "home.html")]
struct HomeTemplate;

#[derive(Template)]
#[template(path = "partial-home.html")]
struct PartialHomeTemplate;

#[derive(Template)]
#[template(path = "about.html")]
struct AboutTemplate;

#[derive(Template)]
#[template(path = "partial-about.html")]
struct PartialAboutTemplate;

#[derive(Template)]
#[template(path = "blog.html")]
struct BlogTemplate;

#[derive(Template)]
#[template(path = "partial-blog.html")]
struct PartialBlogTemplate;

struct HtmlTemplate<T>(T);

impl<T> IntoResponse for HtmlTemplate<T>
where
    T: Template,
{
    fn into_response(self) -> Response {
        match self.0.render() {
            Ok(html) => Html(html).into_response(),
            Err(err) => (
                StatusCode::INTERNAL_SERVER_ERROR,
                format!("Failed to render template. Error: {}", err),
            )
                .into_response(),
        }
    }
}
