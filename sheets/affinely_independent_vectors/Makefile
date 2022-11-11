# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affinely_independent_vectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

affinely_independent_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affinely_independent_vectors.tex

affinely_independent_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex affinely_independent_vectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode affinely_independent_vectors.tex
	../../latexrun affinely_independent_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affinely_independent_vectors.pdf
	open affinely_independent_vectors.pdf

o:
	make open

terms: affinely_independent_vectors.tex
	bbk_sheet -mode ts affinely_independent_vectors.tex

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
	rm -f affinely_independent_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile