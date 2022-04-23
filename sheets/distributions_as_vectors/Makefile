# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distributions_as_vectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

distributions_as_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distributions_as_vectors.tex

distributions_as_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex distributions_as_vectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distributions_as_vectors.tex
	../../latexrun distributions_as_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distributions_as_vectors.pdf
	open distributions_as_vectors.pdf

o:
	make open

terms: distributions_as_vectors.tex
	bbk_sheet -mode ts distributions_as_vectors.tex

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
	rm -f distributions_as_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile