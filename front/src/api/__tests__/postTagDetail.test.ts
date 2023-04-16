import { rest } from 'msw';
import { setupServer } from 'msw/node';
import type Tag from '../../types/tag';
import postTagDetail from '../postTagDetail';

const server = setupServer(
  rest.post<Tag>('http://localhost:5173/tag', async (req, res, ctx) => {
    const tag: Tag = await req.json();
    if (tag.name === 'error') {
      return await res(
        ctx.status(500),
        ctx.json({ message: 'Internal Server Error' })
      );
    }

    return await res(
      ctx.status(200),
      ctx.json({
        ...tag,
        id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
      })
    );
  })
);

beforeAll(() => {
  server.listen();
});
afterEach(() => {
  server.resetHandlers();
});
afterAll(() => {
  server.close();
});

describe('postTagDetail', () => {
  it('should return a tag with generated ID when successfully created', async () => {
    const inputTag: Tag = {
      id: '',
      name: 'testName',
      nameReading: 'testNameReading',
    };

    const expectedTag: Tag = {
      id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
      name: 'testName',
      nameReading: 'testNameReading',
    };

    const result = await postTagDetail(inputTag);

    expect(result).toEqual(expectedTag);
  });

  it('should throw an error when API call fails', async () => {
    const inputTag: Tag = {
      id: '',
      name: 'error',
      nameReading: 'testNameReading',
    };

    await expect(postTagDetail(inputTag)).rejects.toThrow(
      'Request failed with status code 500'
    );
  });
});
