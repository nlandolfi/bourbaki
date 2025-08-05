# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make group_embeddings.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

group_embeddings.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > group_embeddings.tex

group_embeddings.pdf: ../../*.tex ../../trademark.pdf *.tex group_embeddings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode group_embeddings.tex
	../../latexrun group_embeddings.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: group_embeddings.pdf
	open group_embeddings.pdf

o:
	make open

terms: group_embeddings.tex
	bbk sheet -mode ts group_embeddings.tex

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
	rm -f group_embeddings.*
	make reset
	bbk sheet -mode mk > Makefile