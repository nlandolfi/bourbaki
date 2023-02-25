# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_random_variables.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_random_variables.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_random_variables.tex

real_random_variables.pdf: ../../*.tex ../../trademark.pdf *.tex real_random_variables.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_random_variables.tex
	../../latexrun real_random_variables.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_random_variables.pdf
	open real_random_variables.pdf

o:
	make open

terms: real_random_variables.tex
	bbk_sheet -mode ts real_random_variables.tex

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
	rm -f real_random_variables.*
	make reset
	bbk_sheet -mode mk > Makefile