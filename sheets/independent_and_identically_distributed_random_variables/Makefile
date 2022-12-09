# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make independent_and_identically_distributed_random_variables.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

independent_and_identically_distributed_random_variables.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > independent_and_identically_distributed_random_variables.tex

independent_and_identically_distributed_random_variables.pdf: ../../*.tex ../../trademark.pdf *.tex independent_and_identically_distributed_random_variables.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode independent_and_identically_distributed_random_variables.tex
	../../latexrun independent_and_identically_distributed_random_variables.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_and_identically_distributed_random_variables.pdf
	open independent_and_identically_distributed_random_variables.pdf

o:
	make open

terms: independent_and_identically_distributed_random_variables.tex
	bbk_sheet -mode ts independent_and_identically_distributed_random_variables.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f independent_and_identically_distributed_random_variables.*
	make reset
	bbk_sheet -mode mk > Makefile