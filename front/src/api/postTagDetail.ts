import axios from '../lib/axios';
import type Tag from '../types/tag';

const postTagDetail = async (tag: Tag): Promise<Tag> => {
  const response = await axios.post<Tag>('tag', tag);

  return response.data;
};

export default postTagDetail;
