# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make feature_embeddings.pdf

feature_embeddings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > feature_embeddings.tex

feature_embeddings.pdf: ../../*.tex ../../trademark.pdf *.tex feature_embeddings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode feature_embeddings.tex
	../../latexrun feature_embeddings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: feature_embeddings.pdf
	open feature_embeddings.pdf

o:
	make open

terms: feature_embeddings.tex
	bbk_sheet -mode ts feature_embeddings.tex

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
	rm -f feature_embeddings.*
	make reset
	bbk_sheet -mode mk > Makefile