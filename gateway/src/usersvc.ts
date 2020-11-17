import axios, { AxiosResponse } from "axios";

const proto = "http://";
const host = process.env.USERSVC_SERVICE_SERVICE_HOST;
const port = process.env.USERSVC_SERVICE_SERVICE_PORT;

export const getUserByID = async (id: string): Promise<AxiosResponse> => {
  return await axios.request({
    method: "GET",
    url: "/users/" + id,
    baseURL: proto + host + ":" + port,
  });
};

export const createUser = async (
  data: Record<string, unknown>
): Promise<AxiosResponse> => {
  return await axios.request({
    method: "POST",
    url: "/users",
    baseURL: proto + host + ":" + port,
    headers: {
      "Content-Type": "application/json",
    },
    data: JSON.stringify(data),
  });
};
