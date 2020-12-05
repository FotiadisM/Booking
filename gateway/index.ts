import express from "express";
import bodyParser from "body-parser";
import usersRouter from "./src/usersvc";
import listingRouter from "./src/listingsvc";

const app = express();
app.use(bodyParser.json());

const port = 8080;

app.get("/", (_, res) => {
  res.send("HELLOOO WORLD!!");
});

app.use("/users/", usersRouter);
app.use("/listings/", listingRouter);

app.listen(port, () => {
  console.log(`service=gateway port=${port}`);
});
