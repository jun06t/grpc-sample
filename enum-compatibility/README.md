# Enum compatibility check
## Premise
Old proto is following.

```proto
message HelloReply {
  string message = 1;
  enum Role {
    NONE     = 0;
    STANDARD = 1;
    PREMIUM  = 2;
    ULTIMATE = 3;
  }
  Role role = 2;
}
```

Then I update like this.
```proto
message HelloReply {
  string message = 1;
  RoleType role = 2;
}

enum RoleType {
  NONE     = 0;
  STANDARD = 1;
  PREMIUM  = 2;
  ULTIMATE = 3;
}
```

## Confirm behaviour
### Run gRPC server
Server depends on new enum proto ``RoleType``.
```
go run server/main.go
```

### Run gRRC Client
Client depends on old enum proto ``HelloReply.Role``.
```
go run client/main.go
```

### Result
Then it returns
```
2024/11/13 08:50:49 Reply:  Hello alice
2024/11/13 08:50:49 Role:  ULTIMATE
```
This means Enum's rename doesn't break compatibility.