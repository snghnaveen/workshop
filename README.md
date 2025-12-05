Test Req-Res flow 
 Controller 
    LongURL
        - Record does not exists
        - Record does exists
        - Bad input
        - 
    ShortURL
        - Record does not exists
        - Record does exists
        - Bad input

    Config
    - DB Connectio Function
      - Inputs validation
  - Src 
    - Code 
    - NewShortCodeGenerator - Sequence till N
      - Benchmarking
    - DB
      - GetFirstURL 
        - Record exists
        - Record does not exits
      - CreateAndGetShortURL
        - Record exists
        - Record does not exits
      - ShortUrlToLongUrl
      - Record exists
        - Record does not exits