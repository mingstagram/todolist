import axios from "axios";

export const getTasksByDate = async (date) => {
  try {
    const response = await axios.get("/tasks?date=" + date);
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};

export const saveTasks = async (taskData) => {
  try {
    const response = await axios.post("/tasks", taskData);
    return response.data;
  } catch (error) {
    console.error("Error creating tasks:", error);
    throw error;
  }
};

export const countTasksByDate = async (date) => {
  try {
    const response = await axios.get(`/tasks/count?date=${date}`);
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};

export const updateCheckedStatus = async (id, isChecked) => {
  try {
    const response = await axios.put("/tasks/checked", {
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
    const response = await axios.delete("/tasks?id=" + id);
    return response.data;
  } catch (error) {
    console.error("Error load todaytasks:", error);
    throw error;
  }
};
