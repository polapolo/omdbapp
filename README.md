Requirements
- Docker

Setup
- `docker-compose up --build`
- Wait until http & grpc served
- Try to hit it

Reset All
- `docker-compose down`
- `docker-compose up --build`

Mysql Client
- Host: `http://localhost:8080/`
- System: `MySQL`
- Username: `root`
- Password: `supersecretpassword`
- Database: `omdb`

HTTP
- Host: `http://localhost:8888`
- Search
    - Endpoint: `http://localhost:8888/search?keyword=batman&page=1`
    - Response:
        ```
        {
            "data": {
                "Search": [
                    {
                        "Title": "Batman Begins",
                        "Year": "2005",
                        "imdbID": "tt0372784",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman v Superman: Dawn of Justice",
                        "Year": "2016",
                        "imdbID": "tt2975590",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman",
                        "Year": "1989",
                        "imdbID": "tt0096895",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman Returns",
                        "Year": "1992",
                        "imdbID": "tt0103776",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman Forever",
                        "Year": "1995",
                        "imdbID": "tt0112462",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman & Robin",
                        "Year": "1997",
                        "imdbID": "tt0118688",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"
                    },
                    {
                        "Title": "The Lego Batman Movie",
                        "Year": "2017",
                        "imdbID": "tt4116284",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman: The Animated Series",
                        "Year": "1992–1995",
                        "imdbID": "tt0103359",
                        "Type": "series",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman: Under the Red Hood",
                        "Year": "2010",
                        "imdbID": "tt1569923",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
                    },
                    {
                        "Title": "Batman: The Dark Knight Returns, Part 1",
                        "Year": "2012",
                        "imdbID": "tt2313197",
                        "Type": "movie",
                        "Poster": "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"
                    }
                ],
                "totalResults": "399",
                "Response": "True",
                "Error": ""
            },
            "error": ""
        }
        ```
- Detail
    - Endpoint: `http://localhost:8888/detail/tt0372784`
    - Response:
        ```
        {
            "data": {
                "Title": "Batman Begins",
                "Year": "2005",
                "Rated": "PG-13",
                "Released": "15 Jun 2005",
                "Runtime": "140 min",
                "Genre": "Action, Adventure",
                "Director": "Christopher Nolan",
                "Writer": "Bob Kane (characters), David S. Goyer (story), Christopher Nolan (screenplay), David S. Goyer (screenplay)",
                "Actors": "Christian Bale, Michael Caine, Liam Neeson, Katie Holmes",
                "Plot": "After training with his mentor, Batman begins his fight to free crime-ridden Gotham City from corruption.",
                "Language": "English, Mandarin",
                "Country": "USA, UK",
                "Awards": "Nominated for 1 Oscar. Another 14 wins & 73 nominations.",
                "Poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
                "Ratings": [
                    {
                        "Source": "Internet Movie Database",
                        "Value": "8.2/10"
                    },
                    {
                        "Source": "Rotten Tomatoes",
                        "Value": "84%"
                    },
                    {
                        "Source": "Metacritic",
                        "Value": "70/100"
                    }
                ],
                "Metascore": "70",
                "imdbRating": "8.2",
                "imdbVotes": "1,288,835",
                "imdbID": "tt0372784",
                "Type": "movie",
                "DVD": "N/A",
                "BoxOffice": "N/A",
                "Production": "Warner Brothers, Di Bonaventura Pictures",
                "Website": "N/A",
                "Response": "True",
                "Error": ""
            },
            "error": ""
        }
        ```

GRPC
- Host: `localhost:9999`
- Search
    - `proto.OMDB.Search`
    - Parameters
        ```
        {
            "keyword": "batman",
            "page": 1
        }
        ```
    - Response
        ```
        {
            "data": {
                "search": [
                {
                    "title": "Batman Begins",
                    "year": "2005",
                    "imdb_id": "tt0372784",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
                },
                {
                    "title": "Batman v Superman: Dawn of Justice",
                    "year": "2016",
                    "imdb_id": "tt2975590",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"
                },
                {
                    "title": "Batman",
                    "year": "1989",
                    "imdb_id": "tt0096895",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BMTYwNjAyODIyMF5BMl5BanBnXkFtZTYwNDMwMDk2._V1_SX300.jpg"
                },
                {
                    "title": "Batman Returns",
                    "year": "1992",
                    "imdb_id": "tt0103776",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BOGZmYzVkMmItM2NiOS00MDI3LWI4ZWQtMTg0YWZkODRkMmViXkEyXkFqcGdeQXVyODY0NzcxNw@@._V1_SX300.jpg"
                },
                {
                    "title": "Batman Forever",
                    "year": "1995",
                    "imdb_id": "tt0112462",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg"
                },
                {
                    "title": "Batman & Robin",
                    "year": "1997",
                    "imdb_id": "tt0118688",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BMGQ5YTM1NmMtYmIxYy00N2VmLWJhZTYtN2EwYTY3MWFhOTczXkEyXkFqcGdeQXVyNTA2NTI0MTY@._V1_SX300.jpg"
                },
                {
                    "title": "The Lego Batman Movie",
                    "year": "2017",
                    "imdb_id": "tt4116284",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg"
                },
                {
                    "title": "Batman: The Animated Series",
                    "year": "1992–1995",
                    "imdb_id": "tt0103359",
                    "type": "series",
                    "poster": "https://m.media-amazon.com/images/M/MV5BOTM3MTRkZjQtYjBkMy00YWE1LTkxOTQtNDQyNGY0YjYzNzAzXkEyXkFqcGdeQXVyOTgwMzk1MTA@._V1_SX300.jpg"
                },
                {
                    "title": "Batman: Under the Red Hood",
                    "year": "2010",
                    "imdb_id": "tt1569923",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
                },
                {
                    "title": "Batman: The Dark Knight Returns, Part 1",
                    "year": "2012",
                    "imdb_id": "tt2313197",
                    "type": "movie",
                    "poster": "https://m.media-amazon.com/images/M/MV5BMzIxMDkxNDM2M15BMl5BanBnXkFtZTcwMDA5ODY1OQ@@._V1_SX300.jpg"
                }
                ],
                "total_results": "399",
                "response": "True",
                "error": ""
            },
            "error": ""
            }
        ```
- Detail
    - `proto.OMDB.MovieDetail`
    - Parameters
        ```
        {
            "id": "tt0372784"
        }
        ```
    - Response
        ```
        {
            "data": {
                "ratings": [
                {
                    "source": "Internet Movie Database",
                    "value": "8.2/10"
                },
                {
                    "source": "Rotten Tomatoes",
                    "value": "84%"
                },
                {
                    "source": "Metacritic",
                    "value": "70/100"
                }
                ],
                "title": "Batman Begins",
                "year": "2005",
                "rated": "PG-13",
                "released": "15 Jun 2005",
                "runtime": "140 min",
                "genre": "Action, Adventure",
                "director": "Christopher Nolan",
                "writer": "Bob Kane (characters), David S. Goyer (story), Christopher Nolan (screenplay), David S. Goyer (screenplay)",
                "actors": "Christian Bale, Michael Caine, Liam Neeson, Katie Holmes",
                "plot": "After training with his mentor, Batman begins his fight to free crime-ridden Gotham City from corruption.",
                "language": "English, Mandarin",
                "country": "USA, UK",
                "awards": "Nominated for 1 Oscar. Another 14 wins & 73 nominations.",
                "poster": "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
                "metascore": "70",
                "imdb_rating": "8.2",
                "imdb_votes": "1,288,835",
                "imdb_id": "tt0372784",
                "type": "movie",
                "dvd": "N/A",
                "box_office": "N/A",
                "production": "Warner Brothers, Di Bonaventura Pictures",
                "website": "N/A",
                "response": "True",
                "error": ""
            },
            "error": ""
        }
        ```