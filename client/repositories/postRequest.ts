import axios from 'axios';
type TProps<T> = {
  url: string;
  formData: T;
};
export const postRequest = async <T, U = {}>({
  url,
  formData,
}: TProps<T>): Promise<U> => {
  const data = await axios
    .post(`http://localhost:8080/${url}`, formData)
    .then((res) => {
      return res.data;
    })
    .catch((error) => {
      console.log(error);
    });
  console.log(data);
  return data;
};
