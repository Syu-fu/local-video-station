import axios from '../lib/axios';
import type Video from '../types/video';

const getVideoDetail = async (id: string): Promise<Video> => {
  const response = await axios.get<Video>(`video/${id}`);

  return response.data;
};

export default getVideoDetail;
