import express from "express";
import axios from "axios";

const proto = "http://";
const host = process.env.LISTINGSVC_SERVICE_SERVICE_HOST;
const port = process.env.LISTINGSVC_SERVICE_SERVICE_PORT;
const url = proto + host + ":" + port;

const router = express.Router();

router.get("/", async (req, res) => {
  try {
    const response = await axios.request({
      method: "GET",
      url: "/users",
      baseURL: url,
    });

    if (response.status === 200) {
      res.send(response.data);
    }
  } catch (err) {
    console.log(err);
    res.statusCode = 500;
    res.json({ error: "Internal Server Error" });
  }
});

router.get("/:id", async (req, res) => {
  try {
    const response = await axios.request({
      method: "GET",
      url: "/users/" + req.params.id,
      baseURL: url,
    });

    if (response.status === 200) {
      res.send(response.data);
    }
  } catch (err) {
    console.log(err);
    res.statusCode = 500;
    res.json({ error: "Internal Server Error" });
  }
});

router.post("/", async (req, res) => {
  try {
    const response = await axios.request({
      method: "POST",
      url: "/users",
      baseURL: url,
      headers: {
        "Content-Type": "application/json",
      },
      data: JSON.stringify(req.body),
    });

    if (response.status === 200) {
      res.send(response.data);
    }
  } catch (err) {
    console.log(err);
    res.statusCode = 500;
    res.json({ error: "Internal Server Error" });
  }
});

export default router;
