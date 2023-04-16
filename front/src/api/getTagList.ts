import axios from '../lib/axios';
import type Tag from '../types/tag';

const getTagList = async (): Promise<Tag[]> => {
  const response = await axios.get<Tag[]>(`tag/list`);

  return response.data;
};

export default getTagList;
