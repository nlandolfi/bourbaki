# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make random_variable_moments.pdf

random_variable_moments.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > random_variable_moments.tex

random_variable_moments.pdf: ../../*.tex ../../trademark.pdf *.tex random_variable_moments.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode random_variable_moments.tex
	../../latexrun random_variable_moments.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_variable_moments.pdf
	open random_variable_moments.pdf

o:
	make open

terms: random_variable_moments.tex
	bbk_sheet -mode ts random_variable_moments.tex

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
	rm -f random_variable_moments.*
	make reset
	bbk_sheet -mode mk > Makefile