# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make recursion_theorem.pdf

recursion_theorem.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > recursion_theorem.tex

recursion_theorem.pdf: ../../*.tex ../../trademark.pdf *.tex recursion_theorem.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode recursion_theorem.tex
	../../latexrun recursion_theorem.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: recursion_theorem.pdf
	open recursion_theorem.pdf

o:
	make open

terms: recursion_theorem.tex
	bbk_sheet -mode ts recursion_theorem.tex

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
	rm -f recursion_theorem.*
	make reset
	bbk_sheet -mode mk > Makefile