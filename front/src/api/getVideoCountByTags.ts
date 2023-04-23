import axios from '../lib/axios';
import type Count from '../types/count';
import type Tag from '../types/tag';

const getVideoCountByTags = async (tag: Tag[]): Promise<Count> => {
  const tagIds = tag.map((tag) => tag.id).join(',');
  const response = await axios.get<Count>(`video/count?tags=${tagIds}`);

  return response.data;
};

export default getVideoCountByTags;
