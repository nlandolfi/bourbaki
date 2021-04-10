# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make surjective_functions.pdf
	make graph.pdf

surjective_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > surjective_functions.tex

surjective_functions.pdf: ../../*.tex ../../trademark.pdf *.tex surjective_functions.tex
	# pdflatex --file-line-error -interaction=nonstopmode surjective_functions.tex
	../../latexrun surjective_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: surjective_functions.pdf
	open surjective_functions.pdf

o:
	make open

terms: surjective_functions.tex
	bbk_sheet -mode ts surjective_functions.tex

spell:
	aspell -c surjective_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f surjective_functions.*
	make reset
	bbk_sheet -mode mk > Makefile