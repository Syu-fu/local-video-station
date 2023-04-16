import { expect } from 'chai';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import getVideoCount from '../getVideoCount';

const server = setupServer(
  rest.get('http://localhost:5173/video/count', async (_, res, ctx) => {
    return await res(
      ctx.json({
        id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
        title: 'How to Make Sushi',
        titleReading: 'How to Make Sushi',
        url: 'https://example.com/how-to-make-sushi',
        thumbnailUrl: 'https://example.com/sushi-thumbnail.jpg',
        tags: [
          {
            id: '74FF4BBA-18A5-4ABC-9164-4987021D411F',
            name: 'Sushi',
            nameReading: 'Sushi',
          },
          {
            id: 'A082A0B6-63DE-4E97-BA30-B29319863D4F',
            name: 'Cooking',
            nameReading: 'Cooking',
          },
        ],
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

describe('getVideoDetail', () => {
  it('fetches the detail of video', async () => {
    const promise = getVideoCount();
    const count = await promise;
    expect(count).to.deep.equal({
      id: '3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4',
      title: 'How to Make Sushi',
      titleReading: 'How to Make Sushi',
      url: 'https://example.com/how-to-make-sushi',
      thumbnailUrl: 'https://example.com/sushi-thumbnail.jpg',
      tags: [
        {
          id: '74FF4BBA-18A5-4ABC-9164-4987021D411F',
          name: 'Sushi',
          nameReading: 'Sushi',
        },
        {
          id: 'A082A0B6-63DE-4E97-BA30-B29319863D4F',
          name: 'Cooking',
          nameReading: 'Cooking',
        },
      ],
    });
  });
});
