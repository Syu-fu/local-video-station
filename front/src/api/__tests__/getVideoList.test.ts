import { expect } from 'chai';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import getVideoList from '../getVideoList';

const server = setupServer(
  rest.get('http://localhost:5173/video/list', async (req, res, ctx) => {
    const page = req.url.searchParams.get('page');
    if (page === '1') {
      return await res(
        ctx.json([
          {
            id: '12345-67890-12345-67890',
            title: 'How to Make Sushi',
            titleReading: 'How to Make Sushi',
            url: 'https://example.com/how-to-make-sushi.mp4',
            thumbnailUrl: 'https://example.com/how-to-make-sushi.jpg',
          },
          {
            id: '09876-54321-09876-54321',
            title: 'Hiking in the Mountains',
            titleReading: 'Hiking in the Mountains',
            url: 'https://example.com/hiking-in-the-mountains.mp4',
            thumbnailUrl: 'https://example.com/hiking-thumbnail.jpg',
          },
        ])
      );
    } else if (page === '2') {
      return await res(
        ctx.json([
          {
            id: '12345-67890-12345-67891',
            title: 'How to Make Sushi2',
            titleReading: 'How to Make Sushi2',
            url: 'https://example.com/how-to-make-sushi2.mp4',
            thumbnailUrl: 'https://example.com/how-to-make-sushi2.jpg',
          },
          {
            id: '09876-54321-09876-54322',
            title: 'Hiking in the Mountains2',
            titleReading: 'Hiking in the Mountains2',
            url: 'https://example.com/hiking-in-the-mountains2.mp4',
            thumbnailUrl: 'https://example.com/hiking-thumbnail2.jpg',
          },
        ])
      );
    } else {
      return await res(ctx.status(404));
    }
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

describe('getVideoList', () => {
  it('fetches the first page of videos', async () => {
    const promise = getVideoList('1');
    const videos = await promise;
    expect(videos).to.deep.equal([
      {
        id: '12345-67890-12345-67890',
        title: 'How to Make Sushi',
        titleReading: 'How to Make Sushi',
        url: 'https://example.com/how-to-make-sushi.mp4',
        thumbnailUrl: 'https://example.com/how-to-make-sushi.jpg',
      },
      {
        id: '09876-54321-09876-54321',
        title: 'Hiking in the Mountains',
        titleReading: 'Hiking in the Mountains',
        url: 'https://example.com/hiking-in-the-mountains.mp4',
        thumbnailUrl: 'https://example.com/hiking-thumbnail.jpg',
      },
    ]);
  });

  it('fetches the second page of videos', async () => {
    const promise = getVideoList('2');
    const videos = await promise;
    expect(videos).to.deep.equal([
      {
        id: '12345-67890-12345-67891',
        title: 'How to Make Sushi2',
        titleReading: 'How to Make Sushi2',
        url: 'https://example.com/how-to-make-sushi2.mp4',
        thumbnailUrl: 'https://example.com/how-to-make-sushi2.jpg',
      },
      {
        id: '09876-54321-09876-54322',
        title: 'Hiking in the Mountains2',
        titleReading: 'Hiking in the Mountains2',
        url: 'https://example.com/hiking-in-the-mountains2.mp4',
        thumbnailUrl: 'https://example.com/hiking-thumbnail2.jpg',
      },
    ]);
  });

  it('throws an error if the page is not found', async () => {
    try {
      await getVideoList('3');
    } catch (error) {
      if (error instanceof Error) {
        expect(error.message).to.include('404');
      }
    }
  });
});
