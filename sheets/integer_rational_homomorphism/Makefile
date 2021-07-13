# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_rational_homomorphism.pdf

integer_rational_homomorphism.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integer_rational_homomorphism.tex

integer_rational_homomorphism.pdf: ../../*.tex ../../trademark.pdf *.tex integer_rational_homomorphism.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_rational_homomorphism.tex
	../../latexrun integer_rational_homomorphism.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_rational_homomorphism.pdf
	open integer_rational_homomorphism.pdf

o:
	make open

terms: integer_rational_homomorphism.tex
	bbk_sheet -mode ts integer_rational_homomorphism.tex

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
	rm -f integer_rational_homomorphism.*
	make reset
	bbk_sheet -mode mk > Makefile