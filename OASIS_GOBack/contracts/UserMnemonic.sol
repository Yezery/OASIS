// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

contract UserMnemonic {
    mapping(address => string) public  userMnemonic;
    mapping(address => string[3]) private authenticationMetaInformation;
    // 404 user no fund
    // 505 error
    // 400 success
    // 500 user exist
    modifier isAuthenticationMetaInformationCorrect(
        string memory Sp1,
        string memory Sp2,
        string memory Sp3,
        address userAddress
    ) {
        require(
            keccak256(abi.encodePacked(Sp1)) ==
                keccak256(
                    abi.encodePacked(
                        authenticationMetaInformation[userAddress][0]
                    )
                ) &&
                keccak256(abi.encodePacked(Sp2)) ==
                keccak256(
                    abi.encodePacked(
                        authenticationMetaInformation[userAddress][1]
                    )
                ) &&
                keccak256(abi.encodePacked(Sp3)) ==
                keccak256(
                    abi.encodePacked(
                        authenticationMetaInformation[userAddress][2]
                    )
                ),
            "505"
        );
        _;
    }

    function setAuthenticationMetaInformation(
        string memory Sp1,
        string memory Sp2,
        string memory Sp3,
        address userAddress
    ) public returns (string memory) {
        require(
            keccak256(abi.encodePacked(userMnemonic[userAddress])) !=
                keccak256(abi.encodePacked("")),
            "404"
        );
        require(
            keccak256(abi.encodePacked(Sp1)) !=
                keccak256("") &&
                keccak256(abi.encodePacked(Sp2)) !=
                keccak256("") &&
                keccak256(abi.encodePacked(Sp3)) !=
                keccak256(""),
            "505"
        );
         require(
                userAddress != address(0),
            "505"
        );
        string[3] memory SpArray;
        SpArray[0] = Sp1;
        SpArray[1] = Sp2;
        SpArray[2] = Sp3;
        authenticationMetaInformation[userAddress] = SpArray;
        return "400";
    }

    function setMnemonic(string memory Mnemonic, address userAddress)
        public
        returns (string memory)
    {
        require(
            keccak256(abi.encodePacked(userMnemonic[userAddress])) ==
                keccak256(abi.encodePacked("")),
            "500"
        );
        require(
            keccak256(abi.encodePacked(Mnemonic)) !=
                keccak256("") &&
                userAddress != address(0),
            "505"
        );
        userMnemonic[userAddress] = Mnemonic;
        return "400";
    }

    function forgetMnemonic(
        string memory Sp1,
        string memory Sp2,
        string memory Sp3,
        address userAddress
    )
        public  
        view
        isAuthenticationMetaInformationCorrect(Sp1, Sp2, Sp3, userAddress)
        returns (bool)
    {
          require(
                userAddress != address(0),
            "505"
        );
        require(
            keccak256(abi.encodePacked(userMnemonic[userAddress])) !=
                keccak256(abi.encodePacked("")),
            "404"
        );

        return true;
    }

    function reSetMnemonic(
        string memory Sp1,
        string memory Sp2,
        string memory Sp3,
        address userAddress,
        string memory NewMnemonic
    )
        public  
        isAuthenticationMetaInformationCorrect(Sp1, Sp2, Sp3, userAddress)
        returns (bool)
    {
          require(
                userAddress != address(0),
            "505"
        );
        require(
            keccak256(abi.encodePacked(userMnemonic[userAddress])) !=
                keccak256(abi.encodePacked("")),
            "404"
        );
        userMnemonic[userAddress] = NewMnemonic;
        return true;
    }

}
