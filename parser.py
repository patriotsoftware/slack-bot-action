import sys

def newResultLine(jobresult):
    
    match(jobresult.split(':')[1]):
        case "success":
            return f"✅ {jobresult.split(':')[0]} Succeeded.\n"
        case "failure":
            return f"❌ {jobresult.split(':')[0]} Failed. \n"
        case _:
            return f"❕ {jobresult.split(':')[0]} Didn't Run. \n"
           
def format_results(inputString):
    print(inputString)
    # jobresult_list = inputString.splitlines()
    # formatted_list = [newResultLine(result) for result in inputString.splitlines()]
    # print (''.join(formatted_list))

format_results(sys.argv)


