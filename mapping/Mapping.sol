pragma solidity 0.5.2;

contract MappingTest {

	string public name;
	address private owner;

	struct Test{
		uint number;
		string phrase;
	}

	mapping (address => Test) public addressTest;
	mapping (bytes32 => bool) public bytesTest;

	event LogSetData(address testAddr, uint num, string saidWhat);
	event LogSetBool(bytes32 value, bool boolVal);

	constructor(string memory _name)
		external
	{
		owner = msg.sender;
		name = _name;
	}

	function getOwner()
		public
		returns (address)
	{
		return owner;
	}

	function setTestData(address _address, uint _number, string memory _phrase)
		public
	{
		require (_address != address(0), "address empty");
		addressTest[_address].number = _number;
		addressTest[_address].phrase = _phrase;
		emit LogSetData(_address, _number, _phrase);
	}

	function setBytesBool(bytes32 value, bool setBool)//might need to check to see what happens when you use string "true"/"false"
		public
	{
		bytes32[value] = setBool;
		emit LogSetBool(value, setBool);
	}
	
}