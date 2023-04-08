import axios from '../lib/axios';
import type Count from '../types/count';

const getVideoCount = async (): Promise<Count> => {
  const response = await axios.get<Count>(`video/count`);

  return response.data;
};

export default getVideoCount;
