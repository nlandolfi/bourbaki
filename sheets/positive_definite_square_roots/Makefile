# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make positive_definite_square_roots.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

positive_definite_square_roots.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > positive_definite_square_roots.tex

positive_definite_square_roots.pdf: ../../*.tex ../../trademark.pdf *.tex positive_definite_square_roots.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode positive_definite_square_roots.tex
	../../latexrun positive_definite_square_roots.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: positive_definite_square_roots.pdf
	open positive_definite_square_roots.pdf

o:
	make open

terms: positive_definite_square_roots.tex
	bbk_sheet -mode ts positive_definite_square_roots.tex

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
	rm -f positive_definite_square_roots.*
	make reset
	bbk_sheet -mode mk > Makefile