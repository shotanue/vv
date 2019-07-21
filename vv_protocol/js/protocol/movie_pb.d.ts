export class ComingSoonMoviesRequest {
  constructor ();
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ComingSoonMoviesRequest;
}

export class ComingSoonMoviesResponse {
  constructor ();
  getSummaries(): {};
  setSummaries(a: {}): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ComingSoonMoviesResponse;
}

export class HotMoviesRequest {
  constructor ();
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => HotMoviesRequest;
}

export class HotMoviesResponse {
  constructor ();
  getSummaries(): {};
  setSummaries(a: {}): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => HotMoviesResponse;
}

export class MovieDetailRequest {
  constructor ();
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => MovieDetailRequest;
}

export class MovieDetailResponse {
  constructor ();
  getDetail(): {};
  setDetail(a: {}): void;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => MovieDetailResponse;
}

