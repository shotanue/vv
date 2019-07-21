/**
 * @fileoverview gRPC-Web generated client stub for protocol
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';
import {
  ComingSoonMoviesRequest,
  ComingSoonMoviesResponse,
  HotMoviesRequest,
  HotMoviesResponse,
  MovieDetailRequest,
  MovieDetailResponse} from './movie_pb';

export class MovieServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: {};
  options_: { [s: string]: {}; };

  constructor (hostname: string,
               credentials: {},
               options: { [s: string]: {}; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoHotMovies = new grpcWeb.AbstractClientBase.MethodInfo(
    HotMoviesResponse,
    (request: HotMoviesRequest) => {
      return request.serializeBinary();
    },
    HotMoviesResponse.deserializeBinary
  );

  hotMovies(
    request: HotMoviesRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: HotMoviesResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/protocol.MovieService/HotMovies',
      request,
      metadata,
      this.methodInfoHotMovies,
      callback);
  }

  methodInfoComingSoonMovies = new grpcWeb.AbstractClientBase.MethodInfo(
    ComingSoonMoviesResponse,
    (request: ComingSoonMoviesRequest) => {
      return request.serializeBinary();
    },
    ComingSoonMoviesResponse.deserializeBinary
  );

  comingSoonMovies(
    request: ComingSoonMoviesRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: ComingSoonMoviesResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/protocol.MovieService/ComingSoonMovies',
      request,
      metadata,
      this.methodInfoComingSoonMovies,
      callback);
  }

  methodInfoMovieDetail = new grpcWeb.AbstractClientBase.MethodInfo(
    MovieDetailResponse,
    (request: MovieDetailRequest) => {
      return request.serializeBinary();
    },
    MovieDetailResponse.deserializeBinary
  );

  movieDetail(
    request: MovieDetailRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: MovieDetailResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/protocol.MovieService/MovieDetail',
      request,
      metadata,
      this.methodInfoMovieDetail,
      callback);
  }

}

