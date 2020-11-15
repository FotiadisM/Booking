import axios, { AxiosResponse } from "axios";

export const getUserByID = async (id: string): Promise<AxiosResponse> => {
  return await axios.request({
    method: "GET",
    url: "/users/" + id,
    baseURL: "http://localhost:8080",
  });
};

export const createUser = async (
  data: Record<string, unknown>
): Promise<AxiosResponse> => {
  return await axios.request({
    method: "POST",
    url: "/users",
    baseURL: "http://localhost:8080",
    headers: {
      "Content-Type": "application/json",
    },
    data: JSON.stringify(data),
  });
};
