# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make formal_languages.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

formal_languages.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > formal_languages.tex

formal_languages.pdf: ../../*.tex ../../trademark.pdf *.tex formal_languages.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode formal_languages.tex
	../../latexrun formal_languages.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: formal_languages.pdf
	open formal_languages.pdf

o:
	make open

terms: formal_languages.tex
	bbk_sheet -mode ts formal_languages.tex

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
	rm -f formal_languages.*
	make reset
	bbk_sheet -mode mk > Makefile