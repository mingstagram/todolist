import axios from "axios";
import axiosInstance from "../config/axiosConfig";

export const getTasksByDate = async (date, userId) => {
  try {
    const response = await axiosInstance.get(
      `/tasks?date=${date}&userId=${userId}`
    );
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};

export const saveTasks = async (taskData) => {
  try {
    const response = await axiosInstance.post("/tasks", taskData);
    return response.data;
  } catch (error) {
    console.error("Error creating tasks:", error);
    throw error;
  }
};

export const countTasksByDate = async (date, userId) => {
  try {
    const response = await axiosInstance.get(
      `/tasks/count?date=${date}&userId=${userId}`
    );
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};

export const updateCheckedStatus = async (id, isChecked) => {
  try {
    const response = await axiosInstance.put("/tasks/checked", {
      id,
      checked: isChecked,
    });
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};

export const deleteTasks = async (id) => {
  try {
    const response = await axiosInstance.delete("/tasks?id=" + id);
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};
