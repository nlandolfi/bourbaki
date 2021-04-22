# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make random_vectors.pdf
	make graph.pdf

random_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > random_vectors.tex

random_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex random_vectors.tex
	# pdflatex --file-line-error -interaction=nonstopmode random_vectors.tex
	../../latexrun random_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_vectors.pdf
	open random_vectors.pdf

o:
	make open

terms: random_vectors.tex
	bbk_sheet -mode ts random_vectors.tex

spell:
	aspell -c random_vectors.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f random_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile