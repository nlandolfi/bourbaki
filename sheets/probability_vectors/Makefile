# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probability_vectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

probability_vectors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > probability_vectors.tex

probability_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex probability_vectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probability_vectors.tex
	../../latexrun probability_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_vectors.pdf
	open probability_vectors.pdf

o:
	make open

terms: probability_vectors.tex
	bbk sheet -mode ts probability_vectors.tex

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
	rm -f probability_vectors.*
	make reset
	bbk sheet -mode mk > Makefile