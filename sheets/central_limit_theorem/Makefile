# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make central_limit_theorem.pdf
	make graph.pdf

central_limit_theorem.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > central_limit_theorem.tex

central_limit_theorem.pdf: ../../*.tex ../../trademark.pdf *.tex central_limit_theorem.tex
	# pdflatex --file-line-error -interaction=nonstopmode central_limit_theorem.tex
	../../latexrun central_limit_theorem.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: central_limit_theorem.pdf
	open central_limit_theorem.pdf

o:
	make open

terms: central_limit_theorem.tex
	bbk_sheet -mode ts central_limit_theorem.tex

spell:
	aspell -c central_limit_theorem.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f central_limit_theorem.*
	make reset
	bbk_sheet -mode mk > Makefile