#[macro_use] extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[get("/hello/<name>")]
fn hello(name: String) -> String {
    format!("Hello, {}!", name)
}

#[get("/number/<number>")]
fn nunmber(number: i32) -> String {
    format!("Hello, {}!", number)
}




#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index, hello, nunmber])
}