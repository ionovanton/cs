#!bin/bash

template="https://progress-bar.dev"
r="README.md"
name1="[AaDS]"
name2="[Compilers]"
name3="[Computer_Architecture]"
name4="[Computer_Networking]"
name5="[Databases]"
name6="[Distributed_Systems]"
name7="[Formal_Language_Math]"
name8="[Linux]"
name9="[Math]"
name10="[Operating_Systems]"
name11="[Profiling]"
name12="[Programming]"


declare -a src_array=("${name1}/${r}" "${name2}/${r}" "${name3}/${r}" "${name4}/${r}" "${name5}/${r}" "${name6}/${r}" "${name7}/${r}" "${name8}/${r}" "${name9}/${r}" "${name10}/${r}" "${name11}/${r}" "${name12}/${r}")
echo "## Progress"

job ()
{
	count=$(cat ${1} | awk 'FNR==2{ print $2 }')
	total=$(cat ${1} | awk 'FNR==2{ print $4 }')
	percent=$(awk "BEGIN { pc=100*${count}/${total}; i=int(pc); print (pc-i<0.5)?i:i+1 }")
	echo $count $total $percent
}

for s in ${src_array[@]}; do
	name=${s%/*}
	value=$(job $s $name)
	count=$(echo ${value} | awk '{ print $1 }')
	total=$(echo ${value} | awk '{ print $2 }')
	percent=$(echo ${value} | awk '{ print $3 }')
	echo "### ${name}"
	echo "![Progress](${template}/${percent}/?title=${count}/${total})"
done
