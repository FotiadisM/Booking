import express from "express";
import bodyParser from "body-parser";
import { createUser, getUserByID } from "./src/usersvc";

const app = express();
app.use(bodyParser.json());

const port = 8070;

app.get("/", (_, res) => {
  res.send("HELLOOO WORLD!!");
});

app.get("/users/:id", async (req, res) => {
  try {
    const response = await getUserByID(req.params.id);

    if (response.status === 200) {
      res.send(response.data);
    }
  } catch (err) {
    console.log(err);
    res.statusCode = 500;
    res.json({ error: "Internal Server Error" });
  }
});

app.post("/users", async (req, res) => {
  try {
    const response = await createUser(req.body);

    if (response.status === 200) {
      res.send(response.data);
    }
  } catch (err) {
    console.log(err);
    res.statusCode = 500;
    res.json({ error: "Internal Server Error" });
  }
});

app.listen(port, () => {
  console.log(`service=gateway port=${port}`);
});
