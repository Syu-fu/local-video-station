import axios from '../lib/axios';
import type Video from '../types/video';

const getVideoList = async (page: string): Promise<Video[]> => {
  const response = await axios.get<Video[]>(`video/list?page=${page}`);

  return response.data;
};

export default getVideoList;
