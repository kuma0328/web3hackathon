import axios from 'axios';

export const getRequest = async (url: string) => {
  const data = await axios
    .get(`http://localhost:8080/${url}`)
    .then((res) => {
      return res.data;
    })
    .catch((error) => {
      throw new Error(error);
    });
  console.log(data);
  return data;
};
